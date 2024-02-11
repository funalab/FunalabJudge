
import React from 'react'
import { useParams } from 'react-router-dom'
import DefaultLayout from '../components/DefaultLayout'

const SubmitDetailsPage: React.FC = () => {
  const { submitId } = useParams()
  return (
    <DefaultLayout>
      <>
        <h2>提出番号 {submitId}</h2>

      </>
    </DefaultLayout>
  )
}

export default SubmitDetailsPage
