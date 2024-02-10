import { Button } from '@chakra-ui/react'
import React from 'react'

const PrevButton: React.FC<Function> = (handlePrev: Function) => {
  return (
    <>
      <Button onClick={() => { handlePrev() }}>Prev</Button>
    </>
  )
}

export default PrevButton
