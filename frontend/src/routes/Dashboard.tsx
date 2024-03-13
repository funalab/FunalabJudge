import { useState, useEffect } from 'react';
import { DefaultLayout } from "../components/DefaultLayout";
import { CardList } from "../components/CardList"
import { axiosClient } from '../providers/AxiosClientProvider';
import { useParams } from 'react-router-dom';
import { Box, Flex, Heading, Progress, Spacer, Icon } from '@chakra-ui/react';
import { problemWithStatus } from '../components/CardList';
import { MdFlag } from "react-icons/md";

const DashboardPage = () => {
  const [data, setData] = useState([]);
  const { userName } = useParams();
  const [progress, setProgress] = useState(0)
  useEffect(() => {
    (async () => {
      try {
        const { data } = await axiosClient.get(`/getProblemList/${userName}`)
        setData(data);
        var done = 0
        var total = 0
        data.forEach((pws: problemWithStatus) => {
          if (pws.Status) { done += 1 }
          total += 1
        });
        setProgress(done * 100 / total)
      }
      catch (error) {
        console.log(error)
      }
    })()
  }, []);

  return (
    <DefaultLayout>
      <Box my={4}>
        <Flex >
          <Heading>Assignments</Heading>
          <Spacer />
          <Icon as={MdFlag} w={7} h={7} mt="20px" />
        </Flex>
        {progress === 100 ? (
          <Progress value={progress} />
        ) : (
          <Progress hasStripe colorScheme='green' value={progress} />
        )}
      </Box>
      <CardList data={data} />
    </DefaultLayout>
  );
};

export default DashboardPage;
