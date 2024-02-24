import React, { useEffect, useState } from 'react'
import { useLocation, useParams } from "react-router-dom";
import DefaultLayout from '../components/DefaultLayout'
import { Divider, Heading, PopoverHeader, Table, TableCaption, TableContainer, Tbody, Tfoot, Th, Thead, Tr } from '@chakra-ui/react'
import SubmissionTableRow, { Result, SubmissionWithStatusProps } from '../components/SubmissionTableRow';
import { axiosClient } from '../providers/AxiosClientProvider';

const ResultQueuePage: React.FC = () => {
  const { userName } = useParams()
  const location = useLocation();
  const [submissionsWithStatus, setSubmissionWithStatus] = useState<SubmissionWithStatusProps[]>([])
  const [files, setFiles] = useState<File[]>([])
  const [submittedDate, setSubmittedDate] = useState<string>('')
  const [haveNotComplete, setHaveNotComplete] = useState<boolean>(false)

  const pushSubmissionWithStatus = (newSubmission: SubmissionWithStatusProps) => {
    const newSubmissionWithStatus = [...submissionsWithStatus];
    newSubmissionWithStatus.push(newSubmission)
    setSubmissionWithStatus(newSubmissionWithStatus);
  };

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
    setSubmittedDate(submittedDate)

    try {
      const namesAndContents = await retrieveNamesAndContents(files)
      const names = namesAndContents[0];
      const contents = namesAndContents[1];
      const response = await axiosClient.post("/compile", {
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
    axiosClient
      .get(`/submissions/${userName}`)
      .then((response) => {
        const { data } = response;
        setSubmissionWithStatus(data)
        const complete = ["AC", "WA", "CE", "TLE"]
        data.map((submission: SubmissionWithStatusProps) => {
          if (!complete.includes(submission.Status)) {
            setHaveNotComplete(true)
          }
        })
      })
      .catch((error) => {
        console.log(error)
        alert("Failed to fetch data from database")
      })
  }, [])

  /*未確定の奴があるなら0.5sずつリクエストを投げてレンダリングをする*/
  useEffect(() => {
    if (haveNotComplete) {
      const sendStatusRequest = () => {
        axiosClient.get(`/submissions/${userName}`)
          .then((response) => {
            const { data } = response;
            setSubmissionWithStatus(data)
            const complete = ["AC", "WA", "CE", "TLE"]
            let completeFlag = true
            data.map((submission: SubmissionWithStatusProps) => {
              if (!complete.includes(submission.Status)) {
                completeFlag = false
              }
            })
            if (completeFlag === true) {
              clearInterval(intervalId)
            }
          })
          .catch((error) => {
            console.log(error)
            alert("Failed to send status request")
          })
      }
      const intervalId = setInterval(sendStatusRequest, 1000)
      return () => clearInterval(intervalId);
    }
  }, [haveNotComplete])

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
                  UserName={submissionWithStatus.Submission.UserName}
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
