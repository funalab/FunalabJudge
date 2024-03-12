import { useState } from 'react'
import { axiosClient } from '../providers/AxiosClientProvider';
import { ProblemWithTestcase } from '../types/DbTypes';

export const changePid2Pname = ({ id }: { id: number }) => {
  const [pName, setProblemName] = useState("")
  axiosClient
    .get(`/getProblem/${id}`)
    .then((response) => {
      const problem: ProblemWithTestcase = response.data
      setProblemName(problem.Name)
    })
  return pName
}
