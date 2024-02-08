import React, { useEffect, useState } from 'react'
import { Divider, Heading, Text, VStack } from '@chakra-ui/react'
import ExecutionConstraints from '../components/ExecutionConstraints'
import InputOutputBox from '../components/InputOutputBox'
import Testcase, { TestcaseProps } from '../components/Testcase'
import axios from 'axios'

const [name, setName] = useState('')
const [executionTime, setExecutionTime] = useState(0)
const [memoryLimit, setMemoryLimit] = useState(0)
const [statement, setStatement] = useState('')
const [problemConstraints, setProblemConstraints] = useState([])
const [inputFormat, setInputFormat] = useState('')
const [outputFormat, setOutputFormat] = useState('')
const [testcases, setTestcases] = useState([])

interface AssignmentPageProps {
  id: string
}

const AssignmentPage: React.FC<AssignmentPageProps> = async ({ id }) => {
  useEffect(() => {
    /*fetch db and set each parameters.*/
    try {
      const { name, executionTime, memoryLimit, statement,
        problemConstraints, inputFormat,
        outputFormat, testcases } = await axios.get("/assignmentInfo/" + id)
    } catch (error) {
      console.log(error)
      alert("Failed to fetch data from database.")
      /*Temporary error handling*/
    }

    setName(name)
    setExecutionTime(executionTime)
    setMemoryLimit(memoryLimit)
    setStatement(statement)
    setProblemConstraints(problemConstraints)
    setInputFormat(inputFormat)
    setOutputFormat(outputFormat)
    setTestcases(testcases)
  }, [id])
  return (
    <>
      <VStack>
        <Heading>
          {name}
        </Heading>
        <Divider />
        <ExecutionConstraints executionTime={executionTime} memoryLimit={memoryLimit} />
        <VStack>
          <Text>問題文</Text>
          <Text>{statement}</Text>
        </VStack>
        <VStack>
          <Text>入力</Text>
          <Text>入力は以下の形式で標準入力から与えられる。</Text>
          <InputOutputBox content={inputFormat} />
        </VStack>
        <VStack>
          <Text>出力</Text>
          <InputOutputBox content={outputFormat} />
        </VStack>
      </VStack>
      <Divider />
      {testcases.map(
        (testcase: TestcaseProps) => <Testcase
          id={id}
          input={testcase.input}
          output={testcase.output}
        />)}
    </>
  )
}

export default AssignmentPage
