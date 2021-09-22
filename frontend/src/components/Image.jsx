import React from 'react';
import {
  Button,
  Box,
  Image as CImage,
  GridItem,
  useClipboard,
  HStack,
} from '@chakra-ui/react';

import Delete from './Delete';

const Image = ({ images, setImages, img }) => {
  const { url } = img;

  // for the copy button
  const { hasCopied, onCopy } = useClipboard(url);

  return (
    <GridItem>
      <Box>
        <CImage
          width="90%"
          height="auto"
          src={url}
          alt="a screenshot or image"
        />
        <HStack pt="2" spacing={2}>
          <Button colorScheme="blue" onClick={onCopy}>
            {hasCopied ? 'Copied' : 'Copy'}
          </Button>
          <a download href={url}>
            <Button colorScheme="green">Download</Button>
          </a>
          <Delete images={images} setImages={setImages} img={img} />
        </HStack>
      </Box>
    </GridItem>
  );
};

export default Image;
