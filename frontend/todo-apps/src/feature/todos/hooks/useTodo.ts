import { zodResolver } from "@hookform/resolvers/zod"
import { useMemo, useState } from "react"
import { useForm } from "react-hook-form"
import { z } from 'zod'
import { createTodo, CreateTodo } from "../api/Todos"
import { toast } from "sonner"

export const todoFormSchema = z.object({
  title: z.string().min(1).max(50),
  description: z.string().min(1).max(255),
  difficulty: z.number().max(3).min(1).default(1),
  priority: z.number().max(3).min(1).default(1),
})

export const useTodo = (props: { onTodoFetch: () => void }) => {
  const [open, setOpen] = useState(false)
  const form = useForm<CreateTodo>({
    resolver: zodResolver(todoFormSchema),
    defaultValues: {
      title: "",
      description: "",
      difficulty: 1,
      priority: 1,
    },
  })

  const difficulty = useMemo(() => form.getValues("difficulty"), [form]);
  const priority = useMemo(() => form.getValues("priority"), [form]);

  async function onSubmit(todo: CreateTodo) {
    // Do something with the form values.
    // ✅ This will be type-safe and validated.
    try {
      await createTodo(todo)
      toast.success("Create Todo Success")
      props.onTodoFetch()
      setOpen(false)
      form.reset()
    } catch (err) {
      toast.error(`Create Todo Error!!!: ${err}`)
    }
  }

  const difficultyLabel = useMemo(() => {
    return difficulty === 1 ? "Easy" : difficulty === 2 ? "Medium" : "Hard";
  }, [difficulty]);

  const priorityLabel = useMemo(() => {
    return priority === 1 ? "Normal" : priority === 2 ? "High!" : "Urgent!!!!";
  }, [priority]);

  return { open, setOpen, form, onSubmit, difficultyLabel, priorityLabel }
}