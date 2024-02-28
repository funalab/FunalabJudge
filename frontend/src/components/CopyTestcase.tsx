import { Button, VStack, HStack, useClipboard, Text, Textarea } from '@chakra-ui/react'
import React, { useEffect } from 'react'

interface CopyTestcaseProps {
  text: string
  content: string
}

/*
 *
 *This component is used by copy functionality.
  This button would be used when copying testcase into user's clipboard.
  Copy functionality would be implemented in the future.
 * 
 * */

const CopyTestcase: React.FC<CopyTestcaseProps> = ({ text, content }) => {
  const { onCopy, value, setValue, hasCopied } = useClipboard("")

  useEffect(() => {
    setValue(content)
  }, [])

  return (
    <>
      <VStack mb={2} mr={3}>
        <HStack spacing={4}>
          <Text fontSize={24} fontWeight={'bold'}>{text}</Text>
          <Button
            onClick={onCopy}
            _hover={{ bg: "blue.300", color: "white", boxShadow: "xl" }}
          >
            {hasCopied ? "Copied!" : "Copy"}</Button>
        </HStack>
        <Textarea
          value={value}
          mr={2}
          readOnly={true}
          style={{ resize: 'none' }}
        />
      </VStack>
    </>
  )
}

export default CopyTestcase 
