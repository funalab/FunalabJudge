import React, { useEffect, useState } from 'react'
import DefaultLayout from '../components/DefaultLayout'
import { axiosClient } from '../providers/AxiosClientProvider';
import { Select, Table, Thead, Tbody, Tr, Th, Td, Box } from '@chakra-ui/react';
import { FaMedal } from "react-icons/fa";

interface petitCoderStatus {
  ProblemId: number,
  ProblemName: string,
  PCSubmission: PCSubmission[],
}

interface PCSubmission {
  UserName: string,
  SubmittedDate: string,
}

const PetitCoderPage: React.FC = () => {
  const [data, setData] = useState<petitCoderStatus[]>([])
  const [selectedProblemIndex, setSelectedProblemIndex] = useState<number | null>(null);
  const handleSelectChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const index = parseInt(event.target.value);
    setSelectedProblemIndex(index);
  };

  useEffect(() => {
    (async () => {
      try {
        const { data } = await axiosClient.get(`/getPetitCoderStatus`)
        setData(data);
        setSelectedProblemIndex(0);
      }
      catch (error) {
        console.log(error)
      }
    })()
  }, []);
  // problemNameでradio box
  // 各問題の順位表
  // カラム: 順位, 名前, 提出日時
  // closeDateを過ぎたら, 名前をresultページのリンクにする(その問題に絞り込み設定)
  return (
    <DefaultLayout>
    <Box p={5} my={5} >
      <p className='pb-5 font-bold text-2xl' >
        問題を選択
      </p>
      <Select
        // value={data[selectedProblemIndex].ProblemName}
        onChange={handleSelectChange}
        mb={5}
        color={'blue.500'}
        fontWeight={'bold'}
      >
        {data.map((status, index) => (
          <option key={index} value={index}>
            {status.ProblemName}
          </option>
        ))}
      </Select>
      {selectedProblemIndex !== null && (
          <Table>
            <Thead>
              <Tr>
                <Th>順位</Th>
                <Th>名前</Th>
                <Th>提出日時</Th>
              </Tr>
            </Thead>
            <Tbody>
              {data[selectedProblemIndex].PCSubmission.map((submission, index) => (
                <Tr key={index}>
                  {index === 0 ? <Td fontWeight={'bold'} >{FaMedal}{index+1}</Td> : <Td fontWeight={'bold'} >{index+1}</Td>}
                  <Td>{submission.UserName}</Td>
                  <Td>{new Date(submission.SubmittedDate).toLocaleString()}</Td>
                </Tr>
              ))}
            </Tbody>
          </Table>
      )}
    </Box>
    </DefaultLayout>
  );
};

export default PetitCoderPage
