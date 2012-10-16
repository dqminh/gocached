require 'dalli'

describe "Integration Spec" do
  before do
    server = File.join(File.expand_path(File.dirname(File.dirname(__FILE__))), 'main.go')
    puts "starting server"
    @pid = Process.spawn "go run #{server}"
    puts @pid
    sleep 0.5
    @cache = Dalli::Client.new('localhost:12345')
  end

  after do
    puts "shutting down"
    Process.kill "INT", @pid
    Process.wait @pid
  end

  it 'sets the value' do
    @cache.set 'test', 'value'
    value = @cache.get 'test'
    value.should eq('value')
  end
end
