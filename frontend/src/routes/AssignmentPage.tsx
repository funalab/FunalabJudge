import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Divider, Heading, Text, VStack } from "@chakra-ui/react";
import ExecutionConstraints from "../components/ExecutionConstraints";
import InputOutputBox from "../components/InputOutputBox";
import Testcase, { TestcaseProps } from "../components/Testcase";
import DefaultLayout from "../components/DefaultLayout";
import { axiosClient } from "../providers/AxiosClientProvider";

const AssignmentPage = () => {
  const { problemId } = useParams<string>();
  const [name, setName] = useState("");
  const [executionTime, setExecutionTime] = useState(0);
  const [memoryLimit, setMemoryLimit] = useState(0);
  const [statement, setStatement] = useState("");
  const [problemConstraints, setProblemConstraints] = useState([]);
  const [inputFormat, setInputFormat] = useState("");
  const [outputFormat, setOutputFormat] = useState("");
  const [testcases, setTestcases] = useState([]);
  useEffect(() => {
    axiosClient
      .get(`/assignmentInfo/${problemId}`)
      .then((response) => {
        const { data } = response;
        console.log(data)
        setName(data.Name);
        setExecutionTime(data.ExTime);
        setMemoryLimit(data.MemLim);
        setStatement(data.Statement);
        setProblemConstraints(data.PrbConst);
        setInputFormat(data.InputFmt);
        setOutputFormat(data.OutputFmt);
        setTestcases(data.Testcases);
      })
      .catch(() => {
        /*Temporary error handling*/
        console.log("error");
        alert("Failed to fetch data from database.");
      });
  }, [problemId]);

  return (
    <>
      <DefaultLayout>
        <VStack>
          <Heading>{name}</Heading>
          <Divider />
          <ExecutionConstraints
            executionTime={executionTime}
            memoryLimit={memoryLimit}
          />
          <VStack>
            <Text>問題文</Text>
            <Text>{statement}</Text>
          </VStack>
          <VStack>
            <Text>制約</Text>
            <Text>{problemConstraints}</Text>
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
        {testcases.map((testcase: TestcaseProps, index: number) => (
          <Testcase
            id={String(index + 1)}
            InputFileContent={testcase.InputFileContent}
            OutputFileContent={testcase.OutputFileContent}
          />
        ))}
      </DefaultLayout>
    </>
  );
};

export default AssignmentPage;
