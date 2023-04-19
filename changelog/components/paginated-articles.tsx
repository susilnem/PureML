import { ReactNode } from "react";
import Link from "next/link";
import Head from "next/head";
import { defaultPx } from "lib/utils/default-container-px";
import Navbar from "components/core/navbar";
import Footer from "components/core/footer";
import { Box, Button, Container, Divider, Heading, HStack, Text, VStack } from "@chakra-ui/react";

export interface PaginatedArticlesProps {
  page: number;
  children: ReactNode;
  articlesLenght?: number;
  ARTICLES_PER_PAGE?: number;
}

export const PaginatedArticles = ({
  page,
  children,
  articlesLenght,
  ARTICLES_PER_PAGE,
}: PaginatedArticlesProps) => {
  const metaTitle = `${page > 0 ? `Page ${page} -` : ""} PureML Changelog`;

  return (
    <>
      <Head>
        <title>{metaTitle}</title>
        <link rel="icon" href="/favicon.ico" />
        <meta name="title" content={metaTitle} />
        <meta name="description" content="Discover new updates and improvements to PureML." />
        <meta name="image" content="https://changelog.pureml.com/PureMLMetaTagImg.svg" />
        <meta property="og:type" content="website" />
        <meta property="og:url" content="https://changelog.pureml.com" />
        <meta property="og:title" content={metaTitle} />
        <meta
          property="og:description"
          content="Discover new updates and improvements to PureML."
        />
        <meta property="og:image" content="https://changelog.pureml.com/PureMLMetaTagImg.svg" />
        <meta name="twitter:card" content="summary_large_image" />
        <meta name="twitter:url" content="https://changelog.pureml.com" />
        <meta name="twitter:title" content={metaTitle} />
        <meta
          name="twitter:description"
          content="Discover new updates and improvements to PureML."
        />
        <meta name="twitter:image" content="https://changelog.pureml.com/PureMLMetaTagImg.svg" />
      </Head>
      <Navbar />
      <Box w="full" maxW="100vw" overflow="hidden" zIndex="docked">
        <Container maxW="landingMax" px={defaultPx(32)} mt={[86, 86, 140]}>
          <VStack>
            <Heading as="h1" fontSize={["5xl"]} color="black">
              Changelog
            </Heading>
            <Text fontSize="xl" color="gray.700">
              How PureML gets better, every week.
            </Text>
          </VStack>
          <Divider my={16} />
          <VStack spacing={16} divider={<Divider />}>
            {children}
            <VStack align={["stretch", "stretch", "center"]}>
              {page === 0 ? (
                <Link href="/page/1">
                  <Button variant="landingOutline" size="landingLg">
                    Load more
                  </Button>
                </Link>
              ) : (
                <HStack justifyContent="center" spacing={4}>
                  {page > 0 && (
                    <Link href={page === 1 ? "/" : `/page/${page - 1}`}>
                      <Button variant="landingOutline" size="landingLg">
                        Previous page
                      </Button>
                    </Link>
                  )}
                  {articlesLenght / ARTICLES_PER_PAGE - 1 > page ? (
                    <Link href={`/page/${page + 1}`}>
                      <Button variant="landingOutline" size="landingLg">
                        Next page
                      </Button>
                    </Link>
                  ) : null}
                </HStack>
              )}
            </VStack>
          </VStack>
        </Container>
        <Footer />
      </Box>
    </>
  );
};
