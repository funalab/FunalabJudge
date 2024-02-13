import { Text, Stack, Divider, Flex, Button, HStack } from '@chakra-ui/react'
import React, { useState } from 'react'
import SubmitButton from './SubmitButton'
import SubmitFile from './SubmitFile';

export interface SubmitFormProps {
}

const SubmitForm: React.FC<SubmitFormProps> = () => {
  const [inputFields, setInputFields] = useState<JSX.Element[]>([
    <SubmitFile />
  ]);

  const handlePlus = () => {
    const newInputFields = [...inputFields];
    newInputFields.push(
      <SubmitFile />
    );
    setInputFields(newInputFields);
  };

  const handleMinus = () => {
    if (inputFields.length === 1) {
      return;
    }
    const newInputFields = [...inputFields];
    newInputFields.pop();
    setInputFields(newInputFields);
  };

  return (
    <>
      <Divider />
      <HStack>
        <Text fontSize={30} fontWeight={'bold'}>Submit Form</Text>
        <Button onClick={handlePlus}>+</Button>
        <Button onClick={handleMinus}>-</Button>
      </HStack>
      <Stack>
        <Stack>
          {inputFields}
        </Stack>
        <Flex>
          <SubmitButton />
        </Flex>
      </Stack>
    </>
  )
}

export default SubmitForm
