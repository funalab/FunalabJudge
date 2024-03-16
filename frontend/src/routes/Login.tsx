import { useState, FormEvent, FC } from 'react';
import { HttpStatusCode } from 'axios';
import { useLocation, useNavigate } from 'react-router-dom';
import { PasswordField } from '../components/PasswordField';
import {
  Box,
  Button,
  Container,
  FormControl,
  FormLabel,
  Flex,
  Heading,
  Input,
  Stack,
  VStack,
} from '@chakra-ui/react'
import { axiosClient } from '../providers/AxiosClientProvider';
import { jwtDecode } from 'jwt-decode';

export interface MyJwtPayload {
  exp: number,
  orig_iat: number,
  joined: Date,
  user: string,
}

export const Login: FC = () => {
  const [userName, setuserName] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();
  const location = useLocation();

  localStorage.removeItem("authUserName");
  localStorage.removeItem("authJoinedDate");
  localStorage.removeItem("authUserExp");

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();

    axiosClient.post("/login", {
      userName: userName,
      password: password,
    })
      .then((response) => {
        if (response.status === HttpStatusCode.Ok) {
          const jwtToken = jwtDecode<MyJwtPayload>(response.data.token);
          localStorage.setItem("authUserName", jwtToken.user);
          localStorage.setItem("authJoinedDate", jwtToken.joined.toString());
          localStorage.setItem("authUserExp", jwtToken.exp.toString());
          if (location.state) {
            navigate(location.state, { replace: true })
          } else {
            navigate(`/dashboard`, { replace: true })
          }
        } else {
          console.error(response.statusText);
          setError('ログイン情報が間違っています。');
        }
      })
      .catch((error) => {
        if (error.response?.status === HttpStatusCode.Unauthorized) {
          console.error(error);
          setError('ログイン情報が間違っています。');
        } else {
          console.error(error);
          setError('通信に失敗しました。');
        }
      });
  };

  return (
    <Flex w="100vw" h="100wh">
      <Container maxW="lg" py={{ base: '12', md: '24' }} px={{ base: '0', sm: '8' }} >
        <Stack spacing="8">
          <Box
            py={{ base: '0', sm: '8' }}
            px={{ base: '4', sm: '10' }}
            bg={{ base: 'transparent', sm: 'bg.surface' }}
            boxShadow={{ base: 'dark-lg', sm: 'dark-lg' }}
            borderRadius={{ base: 'none', sm: 'xl' }}
          >
            <Stack spacing="6">
              <VStack spacing="5">
                <Heading
                  size={{ base: 'xs', md: 'sm' }}
                  pt={3}
                  pb={5}
                >
                  Login to your account
                </Heading>
                <FormControl>
                  <FormLabel htmlFor="email">User Name</FormLabel>
                  <Input id="userName" type="username" value={userName} autoComplete='username' onChange={(e) => setuserName(e.target.value)} />
                </FormControl>
                <FormControl>
                  <FormLabel htmlFor="password">Password</FormLabel>
                  <PasswordField id="password" value={password} onChange={(e) => setPassword(e.target.value)} />
                </FormControl>
              </VStack>
              {error && <div style={{ color: 'red' }}>{error}</div>}
              <Button
                onClick={handleSubmit}
                _hover={{ bg: "blue.300", color: "white", boxShadow: "xl" }}
              >
                Sign in
              </Button>
            </Stack>
          </Box>
        </Stack>
      </Container>
    </Flex>
  );
};

export default Login;
