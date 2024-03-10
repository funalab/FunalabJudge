import { useState, useEffect } from 'react';
import { DefaultLayout } from "../components/DefaultLayout";
import { B3results } from "../components/B3results"
import { axiosClient } from '../providers/AxiosClientProvider';
import { B3StatusProps } from "../components/B3results";

const B3ResultsPage = () => {
  const [data, setData] = useState<B3StatusProps[]>([]);
  useEffect(() => {
    (async () => {
      try {
        const { data } = await axiosClient.get(`/getB3Status`)
        setData(data);
      }
      catch (error) {
        console.log(error)
      }
    })()
  }, []);

  return (
    <DefaultLayout>
      <B3results data={data} />
    </DefaultLayout>
  );
};

export default B3ResultsPage;
