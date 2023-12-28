import { Flex, Image } from "@chakra-ui/react";
import { Box} from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";

export const TopHeader = () => {
  const navigate = useNavigate();

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
      <Box mt="10px" ml="10px">
        <Image src="sample.png" alt="Logo" onClick={() => navigate("/")} />
      </Box>
    </Flex>
  );
};