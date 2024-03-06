import { useEffect } from 'react'
import axios, { HttpStatusCode } from 'axios'
import { useNavigate } from 'react-router-dom'

const SERVER_IP: string = import.meta.env.VITE_PUBLIC_SERVER_IP;
const BACKEND_PORT: string = import.meta.env.VITE_BACKEND_PORT
axios.defaults.baseURL = `http://${SERVER_IP}:${BACKEND_PORT}`;
axios.defaults.withCredentials = true;

export const axiosClient = axios.create({})

export function AxiosClientProvider({ children }: { children: React.ReactElement }) {
  const navigater = useNavigate()

  useEffect(() => {
    const requestInterceptors = axiosClient.interceptors.request.use()
    const responseInterceptor = axiosClient.interceptors.response.use(
      (response) => {
        if (response.status === HttpStatusCode.Ok) {
          const handleRefreshToken = async () => {
            await axios.get("refresh_token").catch(refreshError => {
              console.error(refreshError.response.error);
            })
          }
          handleRefreshToken();
        }
        return response
      },
      (error) => {
        if (error.response?.status) {
          if (error.response.status === HttpStatusCode.Unauthorized) {
            navigater("login")
          } else {
            console.error(error.response.error);
          }
        } else {
          console.error(error);
        }
        return Promise.reject(error)
      }
    )
    // クリーンアップ
    return () => {
      axiosClient.interceptors.request.eject(requestInterceptors)
      axiosClient.interceptors.response.eject(responseInterceptor)
    }
  }, [])

  return (<>{children}</>)
}
