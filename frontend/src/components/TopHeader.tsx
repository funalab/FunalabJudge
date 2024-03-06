import { Flex, Image, Button, Spacer } from "@chakra-ui/react";
import { Box } from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
import { HttpStatusCode } from "axios";
import { axiosClient } from "../providers/AxiosClientProvider";
import { MdLogout } from "react-icons/md";
import LogoImage from "../../images/funalab.png"

export const TopHeader = () => {
  const navigate = useNavigate()
  const loginUser = localStorage.getItem("authUserName")

  const handleLogout = async () => {
    try {
      const { status } = await axiosClient.post("/logout")
      if (status === HttpStatusCode.Ok) {
        localStorage.removeItem("authUserName");
        localStorage.removeItem("authJoinedDate");
        localStorage.removeItem("authUserExp");
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
      <Box ml="10px">
        <Image src={LogoImage} alt="Logo" width={75} />
      </Box>
      <Spacer />
      <Box mt="10px" mr="10px">
        Hello <b>{loginUser}</b> 👋<br></br>
        Welcome to <b>FunalabJudge</b>
      </Box>
      <Box mt="13px" ml="10px">
        <Button leftIcon={<MdLogout />} colorScheme='teal' variant='solid'  onClick={handleLogout} >
          ログアウト
        </Button>
      </Box>
    </Flex>
  );
};
