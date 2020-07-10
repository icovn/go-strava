import React from 'react';
import "../node_modules/gestalt/dist/gestalt.css"
import { Box, Button, Container, Text } from "gestalt";

import queryString from 'querystring'

function App() {
    console.log(window.location.search);
    const parsed = queryString.parse(window.location.search);
    console.log(parsed);

    const handleClick = (e) => {
        const client_id = 50658;
        const redirect_uri = "http://localhost:3000"
        const response_type = 'code'
        const scope = 'activity:read'
        let url = `https://www.strava.com/oauth/authorize?client_id=${client_id}&redirect_uri=${redirect_uri}&response_type=${response_type}&scope=${scope}`
        window.location.assign(url);
    }

  return (
      <Box color="gray" padding={3}>
          <Container>
              <Box color="white" padding={3}>
                  <Text>Centered content</Text>
                  <Button onClick={handleClick} text="Login" inline />
              </Box>
          </Container>
      </Box>
  );
}

export default App;
