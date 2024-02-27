// Function to handle the onerror event of an image
function handleImageError(imageElement) {
  // Check if the image has a data-src attribute

  // If no data-src is available, replace with a fallback image
  // Replace the broken image with a fallback image
  imageElement.src = "/logo.svg";
  imageElement.alt = "Fallback Image";
}

// Selecting all the images
const images = document.querySelectorAll("img");

images.forEach((el) => {
  el.addEventListener("error", () => {
    handleImageError(el);
  });
});
