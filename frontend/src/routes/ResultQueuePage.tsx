import React, { useEffect, useState } from 'react'
import { useLocation, useParams } from "react-router-dom";
import DefaultLayout from '../components/DefaultLayout'
import { Divider, Heading, Table, TableCaption, TableContainer, Tbody, Tfoot, Th, Thead, Tr } from '@chakra-ui/react'
import axios from 'axios';
import SubmissionTableRow, { Result, SubmissionWithStatusProps } from './SubmissionTableRow';

const ResultQueuePage: React.FC = () => {
  const { userId } = useParams()
  const location = useLocation();
  const [submissionsWithStatus, setSubmissionWithStatus] = useState<SubmissionWithStatusProps[]>([])
  const [files, setFiles] = useState<File[]>([])
  const [problemId, setProblemId] = useState<number>(-1)
  const [submittedDate, setSubmittedDate] = useState<string>('')
  const [ready, setReady] = useState(false)

  const pushSubmissionWithStatus = (newSubmission: SubmissionWithStatusProps) => {
    const newSubmissionWithStatus = [...submissionsWithStatus];
    newSubmissionWithStatus.push(newSubmission)
    setSubmissionWithStatus(newSubmissionWithStatus);
  };

  const is_from_navigation = () => {
    return !(location.state === undefined)
  }
  const retrieveNamesAndContents = async (files: File[]) => {
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
  const sendCompileRequest = async () => {
    const files = location.state.files
    const problemId = location.state.problemId;
    const submittedDate = location.state.submittedDate
    setFiles(files)
    setProblemId(problemId)
    setSubmittedDate(submittedDate)

    try {
      const namesAndContents = await retrieveNamesAndContents(files)
      const names = namesAndContents[0];
      const contents = namesAndContents[1];
      const response = await axios.post("/compile", {
        names: names,
        contents: contents,
        problemId: problemId,
        submittedDate: submittedDate
      })
      return response;
    } catch (error) {
      console.log(error);
      return null;
    }
  }


  useEffect(() => {
    axios
      .get(`/submissions/${userId}`)
      .then((response) => {
        const { data } = response;
        setSubmissionWithStatus(data)
        const files = location.state.files;
        const problemId = location.state.problemId;
        const submittedDate = location.state.submittedDate;
        setFiles(files)
        setProblemId(problemId)
        setSubmittedDate(submittedDate)
        setReady(true)
      })
      .catch(() => {
        console.log('error')
        alert("Failed to fetch data from database")
      })
  }, [])

  useEffect(() => {
    if (ready) {
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
        axios
          .get(`/maxSubmissionId`)
          .then(async (maxSubmissionIdResp) => {
            const maxSubmissionId = maxSubmissionIdResp.data.maxSubmissionId;
            sendCompileRequest().then(async resp => {
              if (resp) {
                const data = resp.data;
                const status = resp.status
                if (status === 200) {
                  /* Throw judge request with websocket or ajax*/
                  console.log(data)
                } else {
                  console.log(status)
                }
              } else {
                try {
                  const currentSubmission = {
                    Id: maxSubmissionId,
                    UserId: userId,
                    ProblemId: problemId,
                    SubmittedDate: submittedDate,
                    Results: [] as Result[],
                    Status: "CE"
                  };
                  pushSubmissionWithStatus({
                    Status: currentSubmission.Status,
                    Submission: currentSubmission
                  })
                  /*push into db*/
                  await axios.post("/addSubmission", currentSubmission)
                } catch (error) {
                  console.log(error)
                }
              }
            });
          })
          .catch(() => {
            console.log('error')
          })
      }
    }
  }, [ready]);

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
