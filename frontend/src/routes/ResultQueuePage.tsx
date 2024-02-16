import React, { useEffect, useState } from 'react'
import { useLocation, useParams } from "react-router-dom";
import DefaultLayout from '../components/DefaultLayout'
import { Divider, Heading, Table, TableCaption, TableContainer, Tbody, Tfoot, Th, Thead, Tr } from '@chakra-ui/react'
import axios from 'axios';
import SubmissionTableRow, { SubmissionWithStatusProps } from './SubmissionTableRow';

const ResultQueuePage: React.FC = () => {
  const { userId } = useParams()
  const location = useLocation();
  const [submissionsWithStatus, setSubmissionWithStatus] = useState<SubmissionWithStatusProps[]>([])

  const is_from_navigation = () => {
    return !(location.state === undefined)
  }
  const retrieveNamsAndContents = async (files: File[]) => {
    const names = [];
    const contents = [];

    for (let fi = 0; fi < files.length; fi++) {
      const reader = new FileReader();
      const file = files[fi];
      const name = file.name;

      const content = await new Promise((resolve, reject) => {
        reader.onload = (event) => resolve(event.target!.result);
        reader.onerror = (error) => reject(error);
        reader.readAsText(file);
      });

      contents.push(content);
      names.push(name);
    }

    return [names, contents];
  }
  useEffect(() => {
    /* fetch all submissions that submitted by user whose id is useId */
    if (is_from_navigation()) {
      /* Should be added ongoing-judge queue row.
       * Before this, we should throw post request into backend.
       * Waiting backend responce, the status should be waiting-judge acronym for WJ.
       * */

      /* Logic is here.
       *
       * 1. throw post request to backend/compile endpoint.
       * 2. get compile result, if CE("compile error") has been occured, UI should be changed, and end.
       * 3. If compile was successful, web socket connection start.
       * */

      /*1. Throw post request to compile endpoint */
      const sendCompileRequest = async () => {
        try {
          const files = location.state.files
          const namesAndContents = await retrieveNamsAndContents(files)
          const names = namesAndContents[0];
          const contents = namesAndContents[1];
          console.log(names)
          console.log(contents)
          const response = await axios.post("/compile", {
            names: names,
            contents: contents,
          })
          return response.data;
        } catch (error) {
          console.log(error);
          return null;
        }
      }
      sendCompileRequest().then(data => {
        if (data) {
          /* Judge Compile Error has been occued or not. */
        } else {
          console.log('No responce from compile endpoint.')
        }
      });
    }
    axios
      .get(`/submissions/${userId}`)
      .then((response) => {
        const { data } = response;
        setSubmissionWithStatus(data)
      })
      .catch(() => {
        console.log('error')
        alert("Failed to fetch data from database")
      })
  }, []);

  return (
    <>
      <DefaultLayout>
        <Heading>自分の提出</Heading>
        <Divider />
        <TableContainer>
          <Table variant='simple'>
            <TableCaption>Your Submission Queue</TableCaption>
            <Thead>
              <Tr>
                <Th>提出日時</Th>
                <Th>問題</Th>
                <Th>ユーザ</Th>
                <Th>結果</Th>
              </Tr>
            </Thead>
            <Tbody>
              {/* This section is ongoing-judge submission row. */}

              {/* This section is existing submission list. */}
              {submissionsWithStatus?.map(submissionWithStatus => (
                <SubmissionTableRow
                  Id={submissionWithStatus.Submission.Id}
                  SubmittedDate={submissionWithStatus.Submission.SubmittedDate}
                  ProblemId={submissionWithStatus.Submission.ProblemId}
                  UserId={submissionWithStatus.Submission.UserId}
                  Results={submissionWithStatus.Submission.Results}
                  Status={submissionWithStatus.Status}
                />
              ))}
            </Tbody>
            <Tfoot>
              {/* Nothing */}
            </Tfoot>
          </Table>
        </TableContainer>
      </DefaultLayout>
    </>
  )
}

export default ResultQueuePage
