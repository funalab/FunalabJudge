import { Flex, Image, Button } from "@chakra-ui/react";
import { Box } from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
import { HttpStatusCode } from "axios";
import { axiosClient } from "../providers/AxiosClientProvider";

export const TopHeader = () => {
  const navigate = useNavigate()
  const loginUser = sessionStorage.getItem("authUserName")

  const handleLogout = async () => {
    try {
      const { status } = await axiosClient.post("/logout")
      if (status === HttpStatusCode.Ok) {
        sessionStorage.removeItem("authUserName");
        sessionStorage.removeItem("authUserRole");
        sessionStorage.removeItem("authUserExp");
        navigate("/login");
      } else {
        // gin-jwtのソースコード的に、OK以外を返すことはない?
        alert("正常にログアウトができませんでした。");
      }
    }
    catch (error) {
      console.error(error);
      alert("正常にログアウトができませんでした。");
    }
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
      zIndex={9999}
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
        <Box mt="10px" ml="10px">
          welcome {loginUser}!
        </Box>
      </Flex>
    </Flex>
  );
};
