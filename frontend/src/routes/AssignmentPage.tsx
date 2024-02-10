import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Divider, Heading, Text, VStack } from "@chakra-ui/react";
import ExecutionConstraints from "../components/ExecutionConstraints";
import InputOutputBox from "../components/InputOutputBox";
import Testcase, { TestcaseProps } from "../components/Testcase";
import axios from "axios";
import DefaultLayout from "../components/DefaultLayout";

const AssignmentPage = () => {
  const { id } = useParams<string>();
  const [name, setName] = useState("");
  const [executionTime, setExecutionTime] = useState(0);
  const [memoryLimit, setMemoryLimit] = useState(0);
  const [statement, setStatement] = useState("");
  const [problemConstraints, setProblemConstraints] = useState([]);
  const [inputFormat, setInputFormat] = useState("");
  const [outputFormat, setOutputFormat] = useState("");
  const [testcases, setTestcases] = useState([]);
  useEffect(() => {
    /*fetch db and set each parameters.
     * 
    type ProblemResp struct {
      Pid       int32      `bson:"problemId"`
      Name      string     `bson:"name"`
      ExTime    int32      `bson: "executionTime"`
      MemLim    int32      `bson: "memoryLimit"`
      Statement string     `bson: "statement"`
      PrbConst  string     `bson: "problemConstraints"`
      InputFmt  string     `bson: "inputFormat"`
      OutputFmt string     `bson: "outputFormat"`
      Testcases []Testcase `bson: "testCases"`
    }
    */

    console.log(id);
    axios
      .get(`/assignmentInfo/${id}`)
      .then((response) => {
        const { data } = response;
        setName(data.Name);
        setExecutionTime(data.ExTime);
        setMemoryLimit(data.MemLim);
        setStatement(data.Statement);
        setProblemConstraints(data.PrbConst);
        setInputFormat(data.InputFmt);
        setOutputFormat(data.OutputFmt);
        setTestcases(data.Testcases);
        console.log();
      })
      .catch(() => {
        /*Temporary error handling*/
        console.log("error");
        alert("Failed to fetch data from database.");
      });
  }, [id]);

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
        {testcases.map((testcase: TestcaseProps) => (
          <Testcase
            id={String(id)}
            InputFile={testcase.InputFile}
            OutputFile={testcase.OutputFile}
          />
        ))}
      </DefaultLayout>
    </>
  );
};

export default AssignmentPage;
