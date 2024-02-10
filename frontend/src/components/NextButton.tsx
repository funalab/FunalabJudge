import { Button } from '@chakra-ui/react'
import React from 'react'

const NextButton: React.FC<Function> = (handleNext: Function) => {
  return (
    <>
      <Button onClick={() => { handleNext() }}>Next</Button>
    </>
  )
}

export default NextButton
