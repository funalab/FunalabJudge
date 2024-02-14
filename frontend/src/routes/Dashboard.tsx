import { useState, useEffect } from 'react';
import { DefaultLayout } from "../components/DefaultLayout";
import { CardList } from "../components/CardList"
import { axiosClient } from '../providers/AxiosClientProvider';
import { useParams } from 'react-router-dom';

export const Dashboard = () => {
  const [data, setData] = useState([]);
  const { userName } = useParams();
  useEffect(() => {
    (async () => {
      try {
        const { data } = await axiosClient.get(`/getAssignmentStatus/${userName}`)
        setData(data);
      }
      catch (error) {
        console.log(error)
      }
    })()
  }, []);

  return (
    <DefaultLayout>
      <CardList data={data} />
    </DefaultLayout>
  );
};

export default Dashboard;
