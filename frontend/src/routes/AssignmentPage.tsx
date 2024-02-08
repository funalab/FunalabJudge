import React from 'react'
import { TestcaseProps } from '../components/Testcase'

/*
 * This is the template for each Assignment.
 * This component recieve props as name, statement, constraints, testcases, submitform.
 * 
 * Testcase,Submitform should be made as another component.
 * */

interface AssignmentPageProps {
  name: string,
  statement: string,
  constraints: string,
  testcases: TestcaseProps[],
}


const AssignmentPage: React.FC<AssignmentPageProps> = ({ name, statement, constraints, testcases, submitform }) => {
  return (
    <>
    </>
  )
}

export default AssignmentPage
