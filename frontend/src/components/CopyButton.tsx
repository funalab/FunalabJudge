import { Button } from '@chakra-ui/react'
import React from 'react'

interface CopyButtonProps {
  content: string
}

/*
 *This component is used by copy functionality.
  This button would be used when copying testcase into user's clipboard.
  Copy functionality would be implemented in the future.
 * */

const CopyButton: React.FC/*<CopyButtonProps>*/ = (/*{ content }*/) => {
  return (
    <>
      <Button>Copy</Button>
    </>
  )
}

export default CopyButton
