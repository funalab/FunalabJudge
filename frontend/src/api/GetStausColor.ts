export const getStatusColor = ({ status }: StatusProps) => {
  if (status === 'AC') {
    return "green.400"
  } else {
    return 'orange.400'
  }
}


