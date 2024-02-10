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
      const apiUrl = 'http://localhost:3000/api/assignments';
      try {
        // HTTP GETリクエストの送信
        const response = await axios.get(apiUrl)
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
}
