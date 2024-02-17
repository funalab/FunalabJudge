import { Button } from '@chakra-ui/react'
import React from 'react'
import { useNavigate } from 'react-router-dom'

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
}

const SubmitButton: React.FC<SubmitButtonProps> = ({ selectedFiles }) => {
  const navigate = useNavigate();
  const handleClick = () => {
    /*Confirm whether the files can be fetched.*/
    selectedFiles.map((file: File) => {
      console.log(file);
    })
    /*navigate into submission queue endpoint with files*/
    const navigationLink = "/results/1" /*  should be changed into result queue endpoint., temporary userId == 1*/
    navigate(navigationLink, {
      state: {
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

