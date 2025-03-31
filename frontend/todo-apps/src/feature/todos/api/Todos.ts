import { z } from "zod";
import { todoFormSchema } from "../hooks/useTodo";

export type Todo = {
  id: number
  title: string;
  description: string;
  priority: 1 | 2 | 3;
  difficulty: 1 | 2 | 3;
  completed: boolean;
  created_at: Date;
  updated_at: Date;
}

export type CreateTodo = z.infer<typeof todoFormSchema>
export type UpdateTodo = CreateTodo & {
  completed: boolean;
}

const fetchTodoList = async (): Promise<Todo[] | null> => {
  try {
    const url = process.env.NEXT_PUBLIC_BACKEND_URL;
    if (!url) throw new Error("BACKEND_URL is not defined");

    const response = await fetch(`${url}/todos`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Accept": "application/json",
      },
      cache: "no-store",

    });

    if (!response.ok) {
      throw new Error(`Error: ${response.status} ${response.statusText}`);
    }

    return await response.json();
  } catch (error) {
    console.error("Error fetching todo list:", error);
    return null;
  }
};

const fetchTodo = async (id: number): Promise<Todo | null> => {
  try {
    const url = process.env.NEXT_PUBLIC_BACKEND_URL;
    if (!url) throw new Error("BACKEND_URL is not defined");

    const response = await fetch(`${url}/todos/${id}`);

    if (!response.ok) {
      throw new Error(`Error: ${response.status} ${response.statusText}`);
    }

    return await response.json();
  } catch (error) {
    console.error("Error fetching todo:", error);
    return null;
  }
}

const createTodo = async (todo: CreateTodo): Promise<Todo | null> => {
  try {
    const url = process.env.NEXT_PUBLIC_BACKEND_URL;
    if (!url) throw new Error("BACKEND_URL is not defined");

    const response = await fetch(`${url}/todos`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Accept": "application/json",
      },
      body: JSON.stringify({
        ...todo,
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString(),
        completed: false,
      }),
    });

    if (!response.ok) {
      throw new Error(`Error: ${response.status} ${response.statusText}`);
    }

    return await response.json();
  } catch (error) {
    console.error("Error creating todo:", error);
    return null;
  }

}

const updateTodo = async (todo: CreateTodo, id: number): Promise<Todo | null> => {
  try {
    const url = process.env.NEXT_PUBLIC_BACKEND_URL;
    if (!url) throw new Error("BACKEND_URL is not defined");

    const response = await fetch(`${url}/todos/${id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        "Accept": "application/json",
      },
      body: JSON.stringify({
        ...todo,
      }),
    });

    if (!response.ok) {
      throw new Error(`Error: ${response.status} ${response.statusText}`);
    }

    return await response.json();
  } catch (error) {
    console.error("Error update todo:", error);
    return null;
  }
}

const deleteTodo = async (id: number): Promise<string | null> => {
  try {
    const url = process.env.NEXT_PUBLIC_BACKEND_URL;
    if (!url) throw new Error("BACKEND_URL is not defined");

    const response = await fetch(`${url}/todos/${id}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
        "Accept": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error(`Error: ${response.status} ${response.statusText}`);
    }

    return "Delete Success";
  } catch (error) {
    console.error("Error delete todo:", error);
    return null;
  }
}

export { fetchTodoList, fetchTodo, createTodo, updateTodo, deleteTodo };