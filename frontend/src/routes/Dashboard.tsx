import { useState, useEffect } from 'react';
import { DefaultLayout } from "../components/DefaultLayout";
import { CardList } from "./CardList"
import axios from 'axios';

export const Dashboard = () => {
  const [data, setData] = useState([]);

  // コンポーネントがマウントされた時にHTTPリクエストを送信する
  useEffect(() => {
    // バックエンドサーバーのエンドポイントURLを指定
    const apiUrl = 'http://localhost:3000/api/assignments';

    // HTTP GETリクエストの送信
    axios.get(apiUrl)
      .then(response => {
        // レスポンスを受け取り、stateにセットする
        setData(response.data);
      })
      .catch(error => {
        console.error('Error fetching data:', error);
      });
  }, []);

  return (
    <DefaultLayout>
      <CardList data={data} />
    </DefaultLayout>
  );
}
