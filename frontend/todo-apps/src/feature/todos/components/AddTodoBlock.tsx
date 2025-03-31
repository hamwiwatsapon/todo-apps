import React from 'react'

import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"

import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"

import { Input } from "@/components/ui/input"
import { PlusCircle, XIcon } from 'lucide-react'
import { Textarea } from "@/components/ui/textarea"
import { Button } from '@/components/ui/button'
import { Slider } from '@/components/ui/slider'
import { DialogDescription } from '@radix-ui/react-dialog'
import { useTodo } from '../hooks/useTodo'


export const AddTodoBlock = ({
  onTodoFetch
}: {
  onTodoFetch: () => void;
}) => {
  const { open, setOpen, form, onSubmit, difficultyLabel, priorityLabel } = useTodo({ onTodoFetch });
  return (
    <Dialog open={open}>
      <DialogTrigger>
        <PlusCircle size={50} className="hover:text-[#2a188783] transition-colors 8duration-300" onClick={() => setOpen(true)} />
      </DialogTrigger>
      <DialogContent>
        <DialogHeader className='flex flex-row justify-between items-center'>
          <DialogTitle>Add Todo</DialogTitle>
          <XIcon className={`hover:text-black/50 duration-300 transition-colors`} size={20} onClick={() => setOpen(false)} />
        </DialogHeader>
        <DialogDescription className='hidden'>
          Create todo window
        </DialogDescription>
        <Form {...form}>
          <FormDescription className='hidden'>
            Create todo form
          </FormDescription>
          <form onSubmit={form.handleSubmit(onSubmit)} className='space-y-4'>
            {/* Title Input */}
            <FormField
              control={form.control}
              name="title"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Title</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />


            {/* Description Input */}
            <FormField
              control={form.control}
              name="description"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Description</FormLabel>
                  <FormControl>
                    <Textarea
                      {...field}
                      id="description-textarea" // Unique ID for the Textarea
                      aria-describedby="description-helper-text" // Reference to description text
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />


            {/* Date Picker */}
            {/* <DatePickerDemo
            // onSelect={(date)=> dispatch({type:""})}
            /> */}

            {/* Difficulty Slider */}
            <FormField
              control={form.control}
              name="difficulty"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>
                    Difficulty
                    <span className='text-gray-500'>
                      Current: {difficultyLabel}
                    </span>
                  </FormLabel>
                  <FormControl>
                    <Slider
                      step={1}
                      max={3}
                      min={1}
                      onValueChange={(value) => field.onChange(value[0])}
                      defaultValue={[field.value]}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            {/* Priority Slider */}
            <FormField
              control={form.control}
              name="priority"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>
                    Priority
                    <span className='text-gray-500'>
                      Current: {priorityLabel}
                    </span>
                  </FormLabel>
                  <FormControl>
                    <Slider
                      step={1}
                      max={3}
                      min={1}
                      onValueChange={(value) => field.onChange(value[0])}
                      defaultValue={[(field.value)]}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <div className='flex flex-row gap-2 justify-end'>
              {/* Cancel Button */}
              <Button
                variant="outline"
                type="button"
                onClick={() => setOpen(false)}
              >
                Cancel
              </Button>
              {/* Confirm Button */}
              <Button
                variant="default"
                type="submit"
              >
                Confirm
              </Button>
            </div>
          </form>
        </Form>

      </DialogContent>
    </Dialog>
  );
};
