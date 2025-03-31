"use client";

import dayjs from 'dayjs';
import React, { useCallback } from 'react'
import { DifficultyBlock } from '@/feature/todos/components/DifficultyBlock';
import { ChevronDown, Trash2 } from 'lucide-react';
import { deleteTodo, Todo, UpdateTodo, updateTodo } from '@/feature/todos/api/Todos';
import { toast } from 'sonner';

const priority = (pri: 1 | 2 | 3) => {
  switch (pri) {
    case 1:
      return "bg-blue-400"
    case 2:
      return "bg-yellow-400"
    case 3:
      return "bg-red-400"
  }
}

export const TodoBlock = (props: {
  todo: Todo
  onTodoFetch: () => void
}) => {
  const priorityColor = priority(props.todo.priority)
  const [isOpen, setIsOpen] = React.useState(false);
  const [isHover, setIsHover] = React.useState(false);
  // Define the height class based on the open state
  const height = isOpen ? "max-h-[500px]" : "max-h-0"  // Set max-height to a large value when open and 0 when closed
  const side = isOpen ? "rotate-180" : ""

  const handleMouseEnter = useCallback(() => setIsHover(true), []);
  const handleMouseLeave = useCallback(() => setIsHover(false), []);

  async function handleCheck(check: boolean, id: number) {
    await updateTodo({ ...props.todo, completed: check } as UpdateTodo, id)
    props.onTodoFetch()
    if (check) {
      toast.success(`Checked todo name: "${props.todo.title}"`)
    } else {
      toast.success(`Uncheck todo name: "${props.todo.title}"`)
    }
  }

  async function handleDelete(id: number) {
    await deleteTodo(id)
    props.onTodoFetch()
    toast.success(`Delete todo name: "${props.todo.title}" success`)
  }

  return (
    <div
      className={`bg-[#F4F2FF] p-4 rounded-lg w-full flex flex-col hover:shadow-2xl duration-300 items-start transition-all gap-4`}
      onMouseEnter={handleMouseEnter}
      onMouseLeave={handleMouseLeave}
    >
      <div className='flex flex-row justify-between w-full items-center text-gray-800'>
        <div className='flex flex-row gap-2 items-center'>
          <input
            type="checkbox" className="form-checkbox h-6 w-6 text-[#2B1887] transition-all"
            defaultChecked={props.todo.completed}
            onChange={async (e) => handleCheck(e.target.checked, props.todo.id)}
          />
          <h1 className='font-semibold text-2xl truncate'>{props.todo.title}</h1>
        </div>
        <button
          type='button'
          onClick={() => setIsOpen(!isOpen)}
        >
          <ChevronDown size={25} className={`${side} duration-300 transition-transform`} />
        </button>
      </div>
      <div className='flex flex-row justify-between w-full items-center'>
        <div className='flex flex-row gap-2 w-full'>
          <div className={`${priorityColor} p-1 rounded-lg shadow text-white font-semibold text-base px-3`}>
            {
              dayjs(props.todo.created_at).isSame(dayjs(), "week") ? DayString(dayjs(props.todo.created_at).day()) : dayjs(props.todo.created_at).format("DD/MM/YYYY")
            }
          </div>
          <div className='flex flex-row w-fit'>
            <DifficultyBlock difficult={props.todo.difficulty} priorityColor={priorityColor} />
          </div>
        </div>
        <button className={`text-red-500 hover:text-red-800 transition-color duration-300 ${!isHover && 'text-transparent'}`} onClick={() => handleDelete(props.todo.id)}>
          <Trash2 size={25} />
        </button>
      </div>

      {/* Description section with smooth transition */}
      <div className={`overflow-hidden transition-all duration-300 ${height} text-left text-gray-700 relative w-full`}>
        <p className='font-medium text-base'>{props.todo.description}</p>
        <p className='justify-self-end text-gray-400'>Last Update: {dayjs(props.todo.updated_at).format("D MMMM YYYY HH:mm")}</p>
      </div>
    </div>
  )
}

const DayString = (day: 0 | 1 | 2 | 3 | 4 | 5 | 6) => {
  switch (day) {
    case 0:
      return "Sun"
    case 1:
      return "Mon"
    case 2:
      return "Tue"
    case 3:
      return "Wed"
    case 4:
      return "Thu"
    case 5:
      return "Fri"
    case 6:
      return "Sat"
  }
}
