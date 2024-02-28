import { Button } from '@chakra-ui/react'
import React, { useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { axiosClient } from '../providers/AxiosClientProvider'

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

  const handleClick = async () => {
    /*navigate into submission queue endpoint with files*/
    if (selectedFiles.length === 0) {
      alert('1つ以上のファイルを選択してください')
      return
    }
    await axiosClient.post(`/addSubmission/${userName}`, {
      problemId: problemId,
      submittedDate: new Date(),
      files: selectedFiles,
    }, {
      headers: {
        'content-type': 'multipart/form-data',
      },
    })
    const navigationLink = `/${userName}/results` /*  should be changed into result queue endpoint., temporary userId == 1*/
    navigate(navigationLink)
  }

  return (
    <>
      <Button
        onClick={handleClick}
        _hover={{ bg: "blue.300", color: "white", boxShadow: "xl" }}
      >
        Submit
      </Button >
    </>
  )
}

export default SubmitButton

