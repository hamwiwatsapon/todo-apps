"use client";

import React, { useMemo } from 'react'
import { Clipboard, ClipboardCheck } from 'lucide-react';
import { TodoBlock } from './TodoBlock';
import { AddTodoBlock } from './AddTodoBlock';

import { Todo } from '@/feature/todos/api/Todos';

interface Todos {
  todos: Todo[] | null;
  onTodoFetch: () => void;
}
export const TodoContainer = ({ todos, onTodoFetch }: Todos) => {
  const { todosData, successData } = useMemo(() => {
    if (!todos) return { todosData: [], successData: [] };

    return {
      todosData: todos.filter((todo) => !todo.completed),
      successData: todos.filter((todo) => todo.completed)
    };
  }, [todos])


  return (
    <div className="flex flex-row min-h-screen gap-16 p-6 font-[family-name:var(--font-geist-sans)] bg-white">
      <div className="flex flex-col gap-[32px] sm:items-start bg-transparent w-full min-h-full justify-center">
        <div className="flex flex-col gap-4 bg-[#D5CCFF] rounded-2xl min-w-lg min-h-full p-6 w-full overflow-y-auto">
          <div className="flex flex-row justify-between text-[#2B1887]">
            <h1 className="text-3xl font-bold flex flex-row gap-2 items-center">
              <Clipboard size={50} />
              <span>To-do</span>
            </h1>
            <AddTodoBlock onTodoFetch={onTodoFetch} />
          </div>
          {
            todosData.length > 0 ? todosData.map((todo: Todo) => (
              <TodoBlock
                key={todo.id}
                todo={todo}
                onTodoFetch={onTodoFetch}
              />
            )) :
              <div className="flex flex-row justify-center items-center w-full h-full">
                <h1 className="text-2xl font-bold text-gray-500">No todos available</h1>
              </div>
          }
        </div>
      </div>
      <div className="flex flex-col gap-[32px] sm:items-start bg-transparent w-full min-h-full justify-center">
        <div className="flex flex-col gap-4 bg-[#D5CCFF] rounded-2xl min-w-lg min-h-full p-6 w-full">
          <div className="flex flex-row justify-between text-[#2B1887]">
            <h1 className="text-3xl font-bold flex flex-row gap-2 items-center text-[#2B1887]">
              <ClipboardCheck size={50} />
              <span>Completed</span>
            </h1>
          </div>
          {
            successData.length > 0 ? successData.map((todo: Todo) => (
              <TodoBlock
                key={todo.id}
                todo={todo}
                onTodoFetch={onTodoFetch}
              />
            )) :
              <div className="flex flex-row justify-center items-center w-full h-full">
                <h1 className="text-2xl font-bold text-gray-500">No completed todos available</h1>
              </div>
          }
        </div>
      </div>
    </div>
  )
}