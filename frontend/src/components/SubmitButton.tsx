import { Button } from '@chakra-ui/react'
import React from 'react'
import { useNavigate, useParams } from 'react-router-dom'

/*
 * SubmitButtonProps Interface should handle two submission way.
 * We should handle both git-commit-hash-case and pure-file-case. 
 * This interface would handle the latter case.
 * So if we implement git-commit-hash-case, another interface would be neccesarry.
 * 
 * If authentication would be completed, navigation would work correctly.
 * */
interface SubmitButtonProps {
  selectedFiles: File[]
  problemId: number
}

const SubmitButton: React.FC<SubmitButtonProps> = ({ selectedFiles, problemId }) => {
  const { userName } = useParams()
  const navigate = useNavigate();
  const handleClick = () => {
    /*Confirm whether the files can be fetched.*/
    selectedFiles.map((file: File) => {
      console.log(file);
    })
    /*navigate into submission queue endpoint with files*/
    const navigationLink = `/${userName}/results/${problemId}` /*  should be changed into result queue endpoint., temporary userId == 1*/
    /*POSTでDBにサブミットの情報をpushする*/
    /*一旦WJでpushして、アップデートしていく。結果を非同期で画面を更新していく*/
    navigate(navigationLink, {
      state: {
        problemId: problemId,
        submittedDate: new Date(),
        files: selectedFiles,
        fromNavigation: true
      }
    })
  }

  return (
    <>
      <Button onClick={handleClick}>
        Submit
      </Button >
    </>
  )
}

export default SubmitButton

