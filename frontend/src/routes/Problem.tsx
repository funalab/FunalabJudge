import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Divider, Heading, Stack, Text, VStack } from "@chakra-ui/react";
import ExecutionConstraints from "../components/ExecutionConstraints";
import InputOutputBox from "../components/InputOutputBox";
import TestcaseBox from "../components/TestcaseBox";
import DefaultLayout from "../components/DefaultLayout";
import { axiosClient } from "../providers/AxiosClientProvider";
import SubmitForm from "../components/SubmitForm";
import { ProblemWithTestcase, Testcase } from "../types/DbTypes";

const ProblemPage = () => {
  const { problemId } = useParams<string>();
  const [name, setName] = useState("");
  const [executionTime, setExecutionTime] = useState(0);
  const [memoryLimit, setMemoryLimit] = useState(0);
  const [statement, setStatement] = useState("");
  const [problemConstraints, setProblemConstraints] = useState("");
  const [inputFormat, setInputFormat] = useState("");
  const [outputFormat, setOutputFormat] = useState("");
  const [testcases, setTestcases] = useState<Testcase[]>([]);

  useEffect(() => {
    axiosClient
      .get<ProblemWithTestcase>(`/getProblem/${problemId}`)
      .then((response) => {
        const p: ProblemWithTestcase = response.data;
        setName(p.Name);
        setExecutionTime(p.ExecutionTime);
        setMemoryLimit(p.MemoryLimit);
        setStatement(p.Statement);
        setProblemConstraints(p.Constraints);
        setInputFormat(p.InputFmt);
        setOutputFormat(p.OutputFmt);
        setTestcases(p.Testcases);
      })
      .catch(() => {
        console.log("error");
        alert("Failed to fetch data from database.");
      });
  }, [problemId]);

  return (
      <DefaultLayout>
        <VStack>
          <Heading my={3}>{name}</Heading>
          <Divider />
          <Stack my={6}>
            <ExecutionConstraints
              executionTime={executionTime}
              memoryLimit={memoryLimit}
            />
            <Stack mt={4} mb={8}>
              <Text
                fontSize={24}
                fontWeight={'bold'}
              >
                問題文
              </Text>
              <Text>{statement}</Text>
            </Stack>

            <Stack mb={8}>
              <Text
                fontSize={24}
                fontWeight={'bold'}
              >
                制約
              </Text>
              <Text>{problemConstraints}</Text>
            </Stack>
            {/* <Stack mb={8}>
              <Text
                fontSize={24}
                fontWeight={'bold'}
              >
                入力
              </Text>
              <Text>入力は以下の形式で標準入力から与えられる。</Text>
              <InputOutputBox content={inputFormat} />

            </Stack>
            <Text
              fontSize={24}
              fontWeight={'bold'}
            >
              出力
            </Text>
            <Text>出力は以下の形式で標準出力に出力せよ。</Text>
            <InputOutputBox content={outputFormat} /> */}
          </Stack>
          <Divider />
          <Text fontSize={32} fontWeight={'bold'}>Sample Cases</Text>
          {testcases.slice(0, Math.min(3, testcases.length)).map((testcase: Testcase, index: number) => (
            <>
              <TestcaseBox
                TestcaseId={index + 1}
                ArgsFileContent={testcase.ArgsFileContent}
                InputFileContent={testcase.InputFileContent}
                OutputFileContent={testcase.OutputFileContent}
              />
            </>
          ))}
          <SubmitForm problemId={+problemId!} />
        </VStack>
      </DefaultLayout >
  );
};

export default ProblemPage;
