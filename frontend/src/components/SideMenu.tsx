import { Box, Button, Icon } from "@chakra-ui/react";
import {
  MdMessage,
  MdDashboard,
  MdCalendarMonth,
  MdAccountBox,
} from "react-icons/md";
import { useNavigate } from "react-router-dom";

export const SideMenu = () => {
  const sideMenuItems = [
    {
      name: "Dashboard",
      icon: MdDashboard,
      path: "/",
    },
    {
      name: "Account",
      icon: MdAccountBox,
      path: "/account",
    },
    {
      name: "Message",
      icon: MdMessage,
      path: "/message",
    },
    {
      name: "Schedule",
      icon: MdCalendarMonth,
      path: "/schedule",
    },
  ];
  // natigate関数を取得
  const navigate = useNavigate();
  return (
    <Box w="20vw" m="20px">
      {sideMenuItems.map((item) => (
        <label key={item.name}>
          <Box mt="10px" ml="10px">
            <Button variant="ghost" onClick={() => navigate(item.path)}>
              <Icon as={item.icon} w={7} h={7} mr="13px" />
              {item.name}
            </Button>
          </Box>
        </label>
      ))}
    </Box>
  );
};

