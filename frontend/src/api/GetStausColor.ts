export const getStatusColor = ({ status }: StatusProps) => {
  if (status === 'AC') {
    return "green.400"
  } else if (status === 'NS') {
    return "blue.200"
  } else {
    return 'orange.400'
  }
}


