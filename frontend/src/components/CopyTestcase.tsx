import { Button, Flex, Textarea, useClipboard } from '@chakra-ui/react'
import React, { useEffect } from 'react'

interface CopyTestcaseProps {
  content: string
}

/*
 *
 *This component is used by copy functionality.
  This button would be used when copying testcase into user's clipboard.
  Copy functionality would be implemented in the future.
 * 
 * */

const CopyTestcase: React.FC<CopyTestcaseProps> = ({ content }) => {
  const { onCopy, value, setValue, hasCopied } = useClipboard("")

  useEffect(() => {
    setValue(content)
  }, [])

  return (
    <>
      <Flex mb={2}>
        <Textarea
          value={value}
          mr={2}
          readOnly={true}
          style={{ resize: 'none' }}
        />
        <Button onClick={onCopy}>{hasCopied ? "Copied!" : "Copy"}</Button>
      </Flex>
    </>
  )
}

export default CopyTestcase 
