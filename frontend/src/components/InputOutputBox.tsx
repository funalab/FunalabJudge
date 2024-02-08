import { Highlight } from '@chakra-ui/react'
import React from 'react'

interface InputOutputBoxProps {
  content: string;
}

const InputOutputBox: React.FC<InputOutputBoxProps> = ({ content }) => {
  return (
    <>
      <Highlight
        query={'spotlight'}
        styles={{ px: '2', py: '1', bg: 'gray.400' }}
      >
        {content}
      </Highlight>
    </>
  )
}

export default InputOutputBox 
