import { Button, VStack, HStack, useClipboard, Text, Textarea, Select } from '@chakra-ui/react'
import React, { useEffect, ChangeEvent } from 'react'
import { InputFileContent } from '../types/DbTypes'

interface CopyTestcaseMultiProps {
  text: string
  files: InputFileContent[]
}

/*
 *
 *This component is used by copy functionality.
  This button would be used when copying testcase into user's clipboard.
  Copy functionality would be implemented in the future.
 * 
 * */

const CopyTestcaseMulti: React.FC<CopyTestcaseMultiProps> = ({ text, files }) => {
  const { onCopy, value, setValue, hasCopied } = useClipboard("")

  useEffect(() => {
    setValue(files[0].Content)
  }, [])

  // マルチセレクト可能なCopyBox
  const handleSelectFile = (ev: ChangeEvent<HTMLSelectElement>) => {
    setValue(ev.target.value);
  }
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
        <Select
          onChange={handleSelectFile}
          color={'blue.500'}
          fontStyle={'italic'}
          fontWeight={'bold'}
        >
          {files.length > 0 && (
            files.map((file, i) => (
              i === 0 ? 
                <option value={file.Content} selected>
                  {file.FileName}
                </option>
              :
                <option value={file.Content}>
                  {file.FileName}
                </option>
            ))
          )}
        </Select>
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

export default CopyTestcaseMulti;
