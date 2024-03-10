import { Flex } from '@chakra-ui/react'
import { getStatusColor } from '../api/GetStausColor'

const StatusBlock = ({ status, onClick }: StatusProps) => {
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
      onClick={onClick}  // Flex コンポーネントに onClick プロパティを渡す
      cursor={onClick ? 'pointer' : 'default'} // クリック可能な場合はポインターに変更
    >
      {status}
    </Flex>
  )
}

export default StatusBlock
