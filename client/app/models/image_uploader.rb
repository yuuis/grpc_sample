# frozen_string_literal:true

require Rails.root.join('app', 'gen', 'api', 'image', 'uploader', 'image_uploader_pb')
require Rails.root.join('app', 'gen', 'api', 'image', 'uploader', 'image_uploader_services_pb')

class ImageUploader
  include ActiveModel::Model

  # Upload an image
  def self.chunked_upload(file_path)
    reqs = Enumerator.new do |yielder|
      # fist request
      filename = File.basename(file_path)
      file_meta = Image::Uploader::ImageUploadRequest::FileMeta.new(filename: filename)

      puts "sent file: #{filename}"
      yielder << Image::Uploader::ImageUploadRequest.new(file_meta: file_meta)

      # chunked upload request
      File.open(file_path, 'r') do |f|
        while (chunk = f.read(100.kilobytes))
          puts "sent #{chunk.size}"
          yielder << Image::Uploader::ImageUploadRequest.new(data: chunk)
        end
      end
    end
    puts 'upload start'
    res = stub.upload(reqs)

    # return response as hash
    {
      uuid: res.uuid,
      size: res.size,
      content_type: res.content_type,
      filename: res.filename
    }
  end

  def self.stub
    Image::Uploader::ImageUploadService::Stub.new(config_dsn, :this_channel_is_insecure, timeout: 5)
  end

  def self.config_dsn
    "127.0.0.1:50052"
  end
end
