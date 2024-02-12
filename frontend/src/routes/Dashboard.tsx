import { useState, useEffect } from 'react';
import { DefaultLayout } from "../components/DefaultLayout";
import { CardList } from "../components/CardList"
import axios from 'axios';

export const Dashboard = () => {
  const [data, setData] = useState([]);
  useEffect(() => {
    (async () => {
      try {
        const response = await axios.get("/")
        setData(response.data);
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
