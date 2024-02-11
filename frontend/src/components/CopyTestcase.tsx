import { Button, Flex, Textarea, useClipboard } from '@chakra-ui/react'
import React from 'react'

interface CopyTestcaseProps {
  content: string
}

/*
 *This component is used by copy functionality.
  This button would be used when copying testcase into user's clipboard.
  Copy functionality would be implemented in the future.
 * */

const CopyTestcase: React.FC<CopyTestcaseProps> = ({ content }) => {
  const { onCopy, value, setValue, hasCopied } = useClipboard("")
  return (
    <>
      <Flex mb={2}>
        <Textarea
          placeholder={content}
          value={value}
          onChange={(e) => {
            setValue(e.target.value)
          }}
          mr={2}
        />
        <Button onClick={onCopy}>{hasCopied ? "Copied!" : "Copy"}</Button>
      </Flex>
    </>
  )
}

export default CopyTestcase 
