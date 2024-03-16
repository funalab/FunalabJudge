import { useState, FormEvent, FC } from 'react';
import DefaultLayout from '../components/DefaultLayout'
import { axiosClient } from '../providers/AxiosClientProvider';
import { PasswordField } from '../components/PasswordField'
import {
  Box,
  Button,
  Container,
  FormControl,
  FormLabel,
  Heading,
  Stack,
} from '@chakra-ui/react'
import { HttpStatusCode } from 'axios';
import { useNavigate } from 'react-router-dom';

const AccountPage: FC = () => {
  const [exPass, setExPass] = useState("");
  const [newPass, setNewPass] = useState("");
  const navigate = useNavigate();

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    
    axiosClient.post(`/changePassword`, {
      exPass: exPass,
      newPass: newPass,
    })
    .then((response) => {
      if (response.status === HttpStatusCode.Ok) {
        alert("パスワードの変更が完了しました。");
        navigate(`/dashboard`)
      } else {
        console.error(response.statusText);
        alert('パスワードの変更に失敗しました。');
      }
    })
    .catch((error) => {
      if (error.response?.status === HttpStatusCode.BadRequest) {
        console.error(error);
        alert('パスワードが間違っています。');
      } else if (error.response?.status === HttpStatusCode.InternalServerError) {
        console.error(error);
        alert('パスワードの変更に失敗しました。');
      } else {
        console.error(error);
        alert('通信に失敗しました。');
      }
    });
  };

  return (
    <DefaultLayout>
    <Container maxW="lg" py={{ base: '12', md: '24' }} px={{ base: '0', sm: '8' }} >
    <Stack spacing="8">
      <Stack spacing="6">
        <Stack spacing={{ base: '2', md: '3' }} textAlign="center">
          <Heading size={{ base: 'xs', md: 'sm' }}>パスワード変更</Heading>
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
              <FormLabel htmlFor="exPass">現在のパスワード</FormLabel>
              <PasswordField id="exPass" value={exPass} onChange={(e) => setExPass(e.target.value)}/>
            </FormControl>
            <FormControl>
              <FormLabel htmlFor="newPass">新しいパスワード</FormLabel>
              <PasswordField id="newPass" value={newPass} onChange={(e) => setNewPass(e.target.value)}/>
            </FormControl>
          </Stack>
          <Stack spacing="6">
            <Button onClick={handleSubmit}>パスワードを変更する</Button>
          </Stack>
        </Stack>
      </Box>
    </Stack>
  </Container>
    </DefaultLayout>
  )
}

export default AccountPage
