"use client";

import { fetchTodoList, Todo } from "@/feature/todos/api/Todos";
import { TodoContainer } from "@/feature/todos/components/TodoContainer";
import { useEffect, useState } from "react";

export default function Home() {
  const [todos, setTodos] = useState<Todo[] | null>(null);

  async function fetchTodos() {
    const data = await fetchTodoList();
    setTodos(data);
  }

  useEffect(() => {
    fetchTodos();
  }, []);

  return (
    <main>
      <TodoContainer todos={todos} onTodoFetch={fetchTodos} />
      <footer className="row-start-3 flex gap-[24px] flex-wrap items-center justify-center">
      </footer>
    </main>
  );
}

