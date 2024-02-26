import { Flex } from '@chakra-ui/react'
import React from 'react'
import { getStatusColor } from '../api/GetStausColor'

const StatusBlock = ({ status }: StatusProps) => {
  return (
    <Flex
      width={"10"}
      bg={getStatusColor({ status: status })}
      color={"white"}
      fontWeight={'bold'}
      px={3}
      py={1}
      borderRadius={'md'}
      justifyContent={"center"}
    >
      {status}
    </Flex>
  )
}

export default StatusBlock
