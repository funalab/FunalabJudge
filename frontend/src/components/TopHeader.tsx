import { Flex, Image, Button } from "@chakra-ui/react";
import { Box } from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
import axios, { HttpStatusCode } from "axios";

export const TopHeader = () => {
  const navigate = useNavigate()

  const handleLogout = async () => {
    axios.post("http://localhost:3000/logout")
    .then((response) => {
      if (response.status === HttpStatusCode.Ok) {
        navigate("/login");
      } else {
        // gin-jwtのソースコード的に、OK以外を返すことはない
        alert("エラー: 正常にログアウトができませんでした。");
      }
    }).catch((error) => {
      console.error(error);
      alert("正常にログアウトができませんでした。");
    })
  }

  return (
    <Flex
      as="header"
      position="fixed"
      bg="gray.100"
      top={0}
      width="full"
      height="100px"
      shadow="sm"
      py={4}
      px={8}
    >
      <Flex>
        <Box mt="10px" ml="10px">
          <Image src="sample.png" alt="Logo" onClick={() => navigate("/")} />
        </Box>
        <Box mt="10px" ml="10px">
          <Button onClick={handleLogout}>ログアウト</Button>
        </Box>
      </Flex>
    </Flex>
  );
};