import { Button } from '@chakra-ui/react'
import React from 'react'

/*
 * SubmitButtonProps Interface should handle two submission way.
 * We should handle both git-commit-hash-case and pure-file-case. 
 * This interface would handle the latter case.
 * So if we implement git-commit-hash-case, another interface would be neccesarry.
 * */
interface SubmitButtonProps {
}

const SubmitButton: React.FC<SubmitButtonProps> = () => {
  return (
    <>
      <Button>
        Submit
      </Button>
    </>
  )
}

export default SubmitButton

