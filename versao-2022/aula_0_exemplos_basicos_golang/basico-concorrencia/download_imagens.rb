require 'thread'
require 'open-uri'



image_urls = [
      "https://cv-julian-2022.s3.us-west-2.amazonaws.com/certificados/E02D0099.jpg",
    "https://cv-julian-2022.s3.us-west-2.amazonaws.com/certificados/F9EA66D0.jpg",
    "https://cv-julian-2022.s3.us-west-2.amazonaws.com/certificados/EE3EC6C2.jpg",]

image_queue = Queue.new
downloaded_images = []

# Criando threads para download
image_urls.each do |url|
  Thread.new do
    image_data = URI.open(url).read
    image_queue << image_data
  end  
end



# Thread principal para processamento
Thread.new do
  image_urls.size.times do
    downloaded_images << image_queue.pop

  end

  puts "Todas as imagens foram baixadas!"
  # ... processar as imagens em downloaded_images
end
