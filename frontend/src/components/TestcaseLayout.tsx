import { Box } from '@chakra-ui/react'
import React, { ReactNode } from 'react'

interface TestcaseLayoutProps {
  children: ReactNode
}

const TestcaseLayout: React.FC<TestcaseLayoutProps> = ({ children }) => {
  return (
    <Box >
      {children}
    </Box>
  )
}

export default TestcaseLayout

