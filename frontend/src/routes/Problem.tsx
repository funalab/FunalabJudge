import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Divider, Heading, Stack, Text, VStack } from "@chakra-ui/react";
import ExecutionConstraints from "../components/ExecutionConstraints";
import TestcaseBox from "../components/TestcaseBox";
import DefaultLayout from "../components/DefaultLayout";
import { axiosClient } from "../providers/AxiosClientProvider";
import SubmitForm from "../components/SubmitForm";
import { ProblemWithTestcase, Testcase } from "../types/DbTypes";

const ProblemPage = () => {
  const { problemId } = useParams<string>();
  const [pwt, setPwt] = useState<ProblemWithTestcase>()

  useEffect(() => {
    axiosClient
      .get<ProblemWithTestcase>(`/getProblem/${problemId}`)
      .then((response) => {
        setPwt(response.data)
      })
      .catch((error) => {
        console.log(error);
        alert("Failed to fetch data from database.");
      });
  }, [problemId]);

  return (
      <DefaultLayout>
        <VStack>
          <Heading my={3}>{pwt?.Name}</Heading>
          <Divider />
          <Stack my={6}>
            <ExecutionConstraints
              executionTime={pwt?.ExecutionTime}
              memoryLimit={pwt?.MemoryLimit}
            />
            <Stack mt={4} mb={8}>
              <Text
                fontSize={24}
                fontWeight={'bold'}
              >
                問題文
              </Text>
              <Text whiteSpace="pre-line">{pwt?.Statement}</Text>
            </Stack>
            <Stack mb={8}>
              <Text
                fontSize={24}
                fontWeight={'bold'}
              >
                制約
              </Text>
              <Text whiteSpace="pre-line">{pwt?.Constraints}</Text>
            </Stack>
            <Stack mb={8}>
              <Text
                fontSize={24}
                fontWeight={'bold'}
              >
                入力形式
              </Text>
              <Text whiteSpace="pre-line">{pwt?.InputFmt}</Text>
            </Stack>
            <Stack mb={8}>
              <Text
                fontSize={24}
                fontWeight={'bold'}
              >
                出力形式
              </Text>
              <Text whiteSpace="pre-line">{pwt?.OutputFmt}</Text>
            </Stack>
          </Stack>
          <Divider />
          <Text fontSize={32} fontWeight={'bold'}>Sample Cases</Text>
          {pwt?.Testcases.slice(0, Math.min(3, isNaN(pwt?.Testcases.length) ? 0 : pwt?.Testcases.length)).map((testcase: Testcase) => (
            <TestcaseBox {...testcase} />
          ))}
          <SubmitForm problemId={+problemId!} />
        </VStack>
      </DefaultLayout >
  );
};

export default ProblemPage;
