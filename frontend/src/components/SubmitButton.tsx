import { Button } from '@chakra-ui/react'
import React from 'react'
// import { useNavigate } from 'react-router-dom'

/*
 * SubmitButtonProps Interface should handle two submission way.
 * We should handle both git-commit-hash-case and pure-file-case. 
 * This interface would handle the latter case.
 * So if we implement git-commit-hash-case, another interface would be neccesarry.
 * 
 * If authentication would be completed, navigation would work correctly.
 * */
interface SubmitButtonProps {
}

const SubmitButton: React.FC<SubmitButtonProps> = () => {
  // const navigate = useNavigate();
  return (
    <>
      <Button /*onClick={() => navigate(`/submissions/${userId}`) }*/>
        Submit
      </Button >
    </>
  )
}

export default SubmitButton

