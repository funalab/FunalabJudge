import { useState, useEffect } from 'react';
import { DefaultLayout } from "../components/DefaultLayout";
import { CardList } from "../components/CardList"
import axios from 'axios';

export const Dashboard = () => {
  const [data, setData] = useState([]);

  // コンポーネントがマウントされた時にHTTPリクエストを送信する
  useEffect(() => {
    (async () => {
      // バックエンドサーバーのエンドポイントURLを指定
      try {
        // HTTP GETリクエストの送信
        const response = await axios.get("/")
        // レスポンスを受け取り、stateにセットする
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
