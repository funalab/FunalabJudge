import { HStack, Text } from '@chakra-ui/react'
import React from 'react'

/* executionTime's unit is sec, memoryLimit's unit is MB.*/
export interface ExecutionConstraintsProps {
  executionTime: number,
  memoryLimit: number,
}
const ExecutionConstraints: React.FC<ExecutionConstraintsProps> = ({ executionTime, memoryLimit }) => {
  return (
    <>
      <HStack>
        <Text>
          実行時間制限: {executionTime} sec
        </Text>
        <Text>
          メモリ制限: {memoryLimit} MB
        </Text>
      </HStack>
    </>
  )
}

export default ExecutionConstraints
