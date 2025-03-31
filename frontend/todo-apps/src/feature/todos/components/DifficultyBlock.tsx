import React from 'react'

interface Props {
  difficult: 1 | 2 | 3;
  priorityColor: string;
}

export const DifficultyBlock = (props: Props) => {
  const noColor = "bg-gray-300"

  return (
    <div className='flex flex-row gap-2 items-center'>
      {(() => {
        const components = [];
        for (let i = 0; i < 3; i++) {
          components.push(
            <div
              key={i}
              className={`${i + 1 > props.difficult ? noColor : props.priorityColor} w-10 h-5 rounded-bl-full rounded-tr-full`}
            />);
        }
        return components;
      })()}
    </div>
  )
}