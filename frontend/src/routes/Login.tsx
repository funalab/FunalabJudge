import React, { useState, FormEvent } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { PasswordField } from '../components/PasswordField'
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
} from '@chakra-ui/react'

export const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();

    axios.post("http://127.0.0.1:8000/accounts/login/", {
      email: email,
      password: password,
    })
    .then((response) => {
      if (response.data.authorized) {
        localStorage.setItem('email', response.data.email);
        navigate("/");
      } else {
        console.error(error);
        setError('ログイン情報が間違っています。');
      }
    })
    .catch((error) => {
      console.error(error);
      setError('ログインに失敗しました。');
    });
  };

  return (
    <Flex w="100vw" h="100wh">
    <Container maxW="lg" py={{ base: '12', md: '24' }} px={{ base: '0', sm: '8' }} >
    <Stack spacing="8">
      <Stack spacing="6">
        <Stack spacing={{ base: '2', md: '3' }} textAlign="center">
          <Heading size={{ base: 'xs', md: 'sm' }}>Log in to your account</Heading>
        </Stack>
      </Stack>
      <Box
        py={{ base: '0', sm: '8' }}
        px={{ base: '4', sm: '10' }}
        bg={{ base: 'transparent', sm: 'bg.surface' }}
        boxShadow={{ base: 'none', sm: 'md' }}
        borderRadius={{ base: 'none', sm: 'xl' }}
      >
        <Stack spacing="6">
          <Stack spacing="5">
            <FormControl>
              <FormLabel htmlFor="email">Email</FormLabel>
              <Input id="email" type="email" value={email} onChange={(e) => setEmail(e.target.value)}/>
            </FormControl>
            <PasswordField id="password" type="password" value={password} onChange={(e) => setPassword(e.target.value)}/>
          </Stack>
          {error && <div style={{ color: 'red' }}>{error}</div>}
          <Stack spacing="6">
            <Button onClick={handleSubmit}>Sign in</Button>
          </Stack>
        </Stack>
      </Box>
    </Stack>
  </Container>
  </Flex>
  );
};

export default Login;
